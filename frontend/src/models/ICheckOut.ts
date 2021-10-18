import { CustomerInterface } from "./ICustomer";
import { EmployeeInterface } from "./IEmployee";
import { CheckInInterface } from "./ICheckIn";

export interface CheckOutInterface {

    ID: number,

    CustomerID: number,
    Customer: CustomerInterface,

    EmployeeID: number,
    Employee: EmployeeInterface,

    CheckInID: number,
    CheckIn: CheckInInterface,

    CheckOutTime: Date;
   
   }