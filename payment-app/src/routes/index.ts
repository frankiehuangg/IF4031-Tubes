import { cancelPayment, processPayment } from "./../services/invoice/index";
import { Invoice } from "@prisma/client";
import express, { Request, Response } from "express";
import { HTTPRequestError } from "../errors/HTTPRequestError";
import { BadRequestError } from "../errors/BadRequestError";
import {
  ICreateInvoice,
  IReadInvoice,
  IUpdateInvoice,
  createInvoice,
  readAllInvoice,
  readInvoice,
  updateInvoice,
} from "../services/invoice";
import KafkaConfig from "../data-access/kafka.server";
const router = express.Router();

const kafkaConfig = new KafkaConfig();
kafkaConfig.consume("payment", (value: any) => {
  console.log("📨 Receive message: ", value);
});

router.get("/test", async (req: Request, res: Response) => {
  try {
    const { message } = req.body;
    const kafkaConfig = new KafkaConfig();
    const messages = [{ key: "key1", value: "foobar" }];
    kafkaConfig.produce("payment", messages);

    res.status(200).json({
      status: "Ok!",
      message: "Message successfully send!",
    });
  } catch (error) {
    console.log(error);
  }
});

/*
Create new invoice 
*/
router.post("/invoices", async (req: Request, res: Response) => {
  try {
    const { client_id, seat_id }: ICreateInvoice = req.body;
    if (!client_id || !seat_id) {
      throw new BadRequestError();
    }
    const data = await createInvoice({ client_id, seat_id });
    return res.status(200).send({ success: true, message: "", data: data });
  } catch (err) {
    if (err instanceof HTTPRequestError) {
      return res
        .status(err.statusCode)
        .send({ success: false, message: err.message, data: [] });
    } else {
      return res
        .status(500)
        .send({ success: false, message: "Internal server Error!", data: [] });
    }
  }
});

/*
Read all/specific invoice
*/
router.get("/invoices", async (req: Request, res: Response) => {
  try {
    let data;
    if (req.query.invoice_id) {
      const invoice_id =
        typeof req.query.invoice_id == "string"
          ? Number.parseInt(req.query.invoice_id)
          : 0;
      data = await readInvoice({ invoice_id });
    } else {
      data = await readAllInvoice();
    }
    return res.status(200).send({ success: true, message: "", data: data });
  } catch (err) {
    console.log(err);
    if (err instanceof HTTPRequestError) {
      return res
        .status(err.statusCode)
        .send({ success: false, message: err.message, data: [] });
    } else {
      return res
        .status(500)
        .send({ success: false, message: "Internal server Error!", data: [] });
    }
  }
});

router.patch("/invoices", async (req: Request, res: Response) => {
  try {
    const { invoice_id, payment_status }: IUpdateInvoice = req.body;
    if (!invoice_id || !payment_status) {
      throw new BadRequestError();
    }

    let data;
    if (payment_status === "SUCCESS") {
      data = await processPayment({ invoice_id });
    } else {
      data = await updateInvoice({ invoice_id, payment_status });
    }
    return res.status(200).send({ success: true, message: "", data: data });
  } catch (err) {
    if (err instanceof HTTPRequestError) {
      return res
        .status(err.statusCode)
        .send({ success: false, message: err.message, data: [] });
    } else {
      return res
        .status(500)
        .send({ success: false, message: "Internal server Error!", data: [] });
    }
  }
});


/* 
Purposfuly using get so client can _simulate_ payment by clicking the link 
*/
router.get("/invoices/pay/:invoice_id", async (req: Request, res: Response) => { 
  try {
    const { invoice_id } = req.params; // No need for 'as unknown'
    const parsedId = parseInt(invoice_id as string, 10); // Specify the radix as 10
    const data = await processPayment({invoice_id: parsedId});
    return res.status(200).send("Payment success!")
  } catch (err) {
    return res.status(400).send({success: false, message: "cannot process payment" })
  }
});

router.get("/invoices/cancel/:invoice_id" , async (req: Request, res: Response) => { 
  try {
    const { invoice_id } = req.params 
    const parsedId = parseInt(invoice_id as string, 10); // Specify the radix as 10
    const data = await cancelPayment({ invoice_id: parsedId });
    return res.status(200).send("Payment canceled!")
  } catch (err) {
    return res.status(400).send({success: false, message: "cannot process cancel" })
  }

})

export default router;
