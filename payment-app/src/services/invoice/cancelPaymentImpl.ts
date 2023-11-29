import { Status } from "@prisma/client";
import { IProcessPayment, readInvoice, updateInvoice } from "./index";
import KafkaConfig from "../../data-access/kafka.server";
import prisma from "../../data-access/db.server";

export const processPaymentImpl = async ({ invoice_id }: IProcessPayment) => {
  const payment_status: Status = "SUCCESS";

  const invoice = await readInvoice({invoice_id})
  if (invoice.payment_status === "SUCCESS") {
    throw new Error("Already paid!")
  }
  const prio_invoice =  await prisma.invoice.groupBy({
    by: ['invoice_id'],
    where: {
      seat_id: invoice.seat_id
    },
    _min: {
        priority: true
    }
  })

  if (prio_invoice[0].invoice_id !== invoice.invoice_id) {
    throw new Error("Cannot pay, still queueing!")
  }
  const data = await updateInvoice({ invoice_id, payment_status });
  try {
    const kafkaConfig = new KafkaConfig();
    const messages = [{ key: "key1", value: JSON.stringify(data) }];
    kafkaConfig.produce("payment", messages);
  } catch (error) {
    console.log(error);
  }
  return data;
};
