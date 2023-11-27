/*
  Warnings:

  - The primary key for the `Invoice` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `_id` on the `Invoice` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[client_id,seat_id]` on the table `Invoice` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `seat_id` to the `Invoice` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Invoice" DROP CONSTRAINT "Invoice_pkey",
DROP COLUMN "_id",
ADD COLUMN     "invoice_id" SERIAL NOT NULL,
ADD COLUMN     "seat_id" INTEGER NOT NULL,
ADD CONSTRAINT "Invoice_pkey" PRIMARY KEY ("invoice_id");

-- CreateIndex
CREATE UNIQUE INDEX "Invoice_client_id_seat_id_key" ON "Invoice"("client_id", "seat_id");
