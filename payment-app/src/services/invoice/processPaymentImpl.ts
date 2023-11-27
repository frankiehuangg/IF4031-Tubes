import { Status } from "@prisma/client";
import { IProcessPayment, updateInvoice } from "./index";

export const processPaymentImpl = async ({ invoice_id }: IProcessPayment) => {
  if (Math.random() < 0.1) {
    // Simulate external call with 10% failure rate
    throw new Error("Error occured while processing payment!");
  }
  const payment_status: Status = "SUCCESS";
  // TODO: call ticket API to update seat status
  const data = await updateInvoice({ invoice_id, payment_status });
  return data;
};
