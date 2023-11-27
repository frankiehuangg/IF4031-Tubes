import { $Enums, Invoice } from "@prisma/client";
import { createInvoiceImpl } from "./createInvoiceImpl";
import { deleteInvoiceImpl } from "./deleteInvoiceImpl";
import { readAllInvoiceImpl, readInvoiceImpl } from "./readInvoiceImpl";
import { updateInvoiceImpl } from "./updateInvoiceImpl";
import { processPaymentImpl } from "./processPaymentImpl";

export interface ICreateInvoice {
  client_id: number;
  seat_id: number;
}
export const createInvoice = createInvoiceImpl;

export interface IDeleteInvoice {
  invoice_id: number;
}
export const deleteInvoice = deleteInvoiceImpl;

export interface IUpdateInvoice {
  invoice_id: number;
  client_id?: number;
  seat_id?: number;
  payment_status?: $Enums.Status;
  priority?: number;
}
export const updateInvoice = updateInvoiceImpl;

export interface IReadInvoice {
  invoice_id: number;
}
export const readInvoice = readInvoiceImpl;

export const readAllInvoice = readAllInvoiceImpl;

export interface IProcessPayment {
  invoice_id: number;
}
export const processPayment = processPaymentImpl;
