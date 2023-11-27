import { HTTPRequestError } from "./HTTPRequestError";

export class ConflictError extends HTTPRequestError {
  constructor(message = "Theres conflict error on the resource!") {
    super(409, message);
  }
}
