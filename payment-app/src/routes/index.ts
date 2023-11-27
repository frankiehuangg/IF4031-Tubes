import { processPayment } from "./../services/invoice/index";
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
const router = express.Router();

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
    const { invoice_id } = req.query.invoice_id as unknown as IReadInvoice;
    let data;
    if (invoice_id) {
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

export default router;
