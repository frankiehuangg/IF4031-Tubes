import { Invoice, Prisma } from "@prisma/client";
import prisma from "../../data-access/db.server";
import { ConflictError } from "../../errors/ConflictError";
import { ICreateInvoice } from ".";

export const createInvoiceImpl = async ({
  client_id,
  seat_id,
}: ICreateInvoice) => {
  try {
    const result = await prisma.invoice.create({
      data: {
        client_id: client_id,
        seat_id: seat_id,
      },
    });
    return result;
  } catch (error) {
    if (
      error instanceof Prisma.PrismaClientKnownRequestError &&
      error.code === "P2002"
    ) {
      throw new ConflictError(
        `Invoice with Client Id ${client_id} and Seats Id ${seat_id} already exist!`
      );
    }
  }
};
