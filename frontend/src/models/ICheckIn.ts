import { CustomerInterface } from "./ICustomer";
import { EmployeeInterface } from "./IEmployee";

export interface CheckInInterface {

    ID: number,
    Date_time:  Date,
	CustomerID: number,
	Customer:   CustomerInterface,
	EmployeeID: number,
	Employee:   EmployeeInterface,
	Room: string;
   
   }