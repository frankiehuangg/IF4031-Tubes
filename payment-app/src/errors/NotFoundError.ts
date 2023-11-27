import { HTTPRequestError } from "./HTTPRequestError";

export class NotFoundError extends HTTPRequestError {
  constructor(message: string = "Request not found!") {
    super(404, message);
  }
}
