import { CustomerInterface } from "./ICustomer";
import { EmployeeInterface } from "./IEmployee";
import { RoomInterface } from "./IRoom";

export interface CheckInInterface {

    ID: number,
    Date_time:  Date,
	CheckInID: number,
	CheckIn:   CustomerInterface,
	EmployeeID: number,
	Employee:   EmployeeInterface,
	ReserveID: number,
	Reserve: RoomInterface;
   }