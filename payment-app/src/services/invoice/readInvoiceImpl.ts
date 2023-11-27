import { Invoice } from "@prisma/client";
import prisma from "../../data-access/db.server";
import { NotFoundError } from "../../errors/NotFoundError";
import { IReadInvoice } from ".";

export const readInvoiceImpl = async ({ invoice_id }: IReadInvoice) => {
  const result = await prisma.invoice.findFirst({
    where: {
      invoice_id: invoice_id,
    },
  });
  if (!result) {
    throw new NotFoundError(`Invoice with id ${invoice_id} not found!`);
  }
  return result;
};

export const readAllInvoiceImpl = async () => {
  const result = await prisma.invoice.findMany();
  return result;
};
