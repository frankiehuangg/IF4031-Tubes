import prisma from "../../data-access/db.server";
import { NotFoundError } from "../../errors/NotFoundError";
import { IUpdateInvoice } from ".";

export const updateInvoiceImpl = async (params: IUpdateInvoice) => {
  const result = await prisma.invoice.update({
    where: {
      invoice_id: params.invoice_id,
    },
    data: params,
  });
  if (!result) {
    throw new NotFoundError(`Invoice with id ${params.invoice_id} not found!`);
  }
  return result;
};
