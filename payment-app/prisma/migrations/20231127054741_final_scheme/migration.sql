-- CreateEnum
CREATE TYPE "Status" AS ENUM ('PENDING', 'CANCELED', 'SUCCESS');

-- CreateTable
CREATE TABLE "Invoice" (
    "_id" INTEGER NOT NULL,
    "client_id" INTEGER NOT NULL,
    "payment_status" "Status" NOT NULL DEFAULT 'PENDING',
    "priority" SERIAL NOT NULL,

    CONSTRAINT "Invoice_pkey" PRIMARY KEY ("_id")
);
