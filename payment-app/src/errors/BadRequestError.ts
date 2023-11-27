import { HTTPRequestError } from "./HTTPRequestError";

export class BadRequestError extends HTTPRequestError {
  constructor(message: string = "Bad Request, please include are params!") {
    super(400, message);
  }
}
