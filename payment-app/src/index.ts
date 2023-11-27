import * as dotenv from "dotenv";
import express from "express";
import router from "./routes";
// import swaggerUi from "swagger-ui-express"
// import swaggerDocument from "./swagger.json"
const cookieParser = require("cookie-parser");
dotenv.config();

if (!process.env.PORT) {
  process.exit(1);
}

const PORT: number = parseInt(process.env.PORT as string, 10);
const app = express();

app.use(express.json());
app.use(router);
// app.use("/api-docs", swaggerUi.serve, swaggerUi.setup(swaggerDocument))
app.listen(PORT, () => {
  console.log(`Listening on  http://localhost:${PORT}`);
});
