import { IDeleteInvoice } from "./index";
import { Invoice } from "@prisma/client";
import prisma from "../../data-access/db.server";
import { NotFoundError } from "../../errors/NotFoundError";

export const deleteInvoiceImpl = async ({ invoice_id }: IDeleteInvoice) => {
  const result = await prisma.invoice.delete({
    where: {
      invoice_id: invoice_id,
    },
  });
  if (!result) {
    throw new NotFoundError(`Invoice with id ${invoice_id} not found!`);
  }
  return result;
};
