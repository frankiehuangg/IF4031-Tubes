import { Invoice } from "@prisma/client";
import prisma from "../../data-access/db.server";
import { NotFoundError } from "../../errors/NotFoundError";
import { IReadInvoice, IReadOldestInvoiceFromClient } from ".";

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

export const readOldestInvoiceFromClientImpl = async ({ client_id, seat_id }: IReadOldestInvoiceFromClient) => {
  const result = await prisma.invoice.aggregate({
    where: {
      client_id: client_id,
      seat_id: seat_id,
    },
    _min: {
      priority: true
    }
  });
  if (!result) {
    throw new NotFoundError(`Invoice with client ${client_id} and seat ${seat_id} not found!`);
  }
  return result;
};

export const readAllInvoiceImpl = async () => {
  const result = await prisma.invoice.findMany();
  return result;
};

