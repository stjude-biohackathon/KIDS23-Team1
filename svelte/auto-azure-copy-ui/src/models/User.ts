import type { Order } from "./Order";

export interface User {
  id: number;
  name: string;
  department: string;
  orders: Order[];
}
