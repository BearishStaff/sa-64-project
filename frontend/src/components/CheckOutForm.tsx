import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";


import { CustomerInterface } from "../models/ICustomer";
import { EmployeeInterface } from "../models/IEmployee";
import { CheckInInterface } from "../models/ICheckIn";
import { CheckOutInterface } from "../models/ICheckOut";
import { RoomInterface } from "../models/IRoom";

import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function CheckOutCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [checkIns, setCheckIns] = useState<CheckInInterface[]>([]);
  const [employees, setEmployees] = useState<EmployeeInterface[]>([]);
  const [customers, setCustomers] = useState<CustomerInterface[]>([]);
  const [rooms, setRooms] = useState<RoomInterface[]>([]);
  const [checkOut, setCheckOut] = useState<Partial<CheckOutInterface>>(
    {}
  );

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
  },
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof checkOut;
    setCheckOut({
      ...checkOut,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof checkOut;
    const { value } = event.target;
    setCheckOut({ ...checkOut, [id]: value });
  };

  const getCheckIns = async () => {
    fetch(`${apiUrl}/check_ins/room`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setCheckIns(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getCustomers = async () => {
    fetch(`${apiUrl}/customers`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setCustomers(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getEmployees = async () => {
    fetch(`${apiUrl}/employees`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setEmployees(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getRooms = async () => {
    fetch(`${apiUrl}/rooms`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setRooms(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getCheckIns();
    getCustomers();
    getEmployees();
    getRooms();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      CheckInID: convertType(checkOut.CheckInID),
      CustomerID: convertType(checkOut.CustomerID),
      EmployeeID: convertType(checkOut.EmployeeID),
      CheckOutTime: selectedDate,
      Condition: checkOut.Condition ?? "",
    };

    console.log(data)

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/check_outs`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSuccess(true);
        } else {
          setError(true);
        }
      });
  }

  // console.log(checkIns)
  // console.log(customers)
  // console.log(employees)

  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper className={classes.paper}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              Check-Out Form
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Room:</p>
              <Select
                native
                value={checkOut.CheckInID}
                onChange={handleChange}
                inputProps={{name: "CheckInID",}}
              >
                <option aria-label="None" value="">
                  กรอกห้อง
                </option>
                {checkIns.map((item: CheckInInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Room.Roomnumber}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Customer Name:</p>
              <Select
                native
                value={checkOut.CustomerID}
                onChange={handleChange}
                inputProps={{
                  name: "CustomerID",
                }}
              >
                <option aria-label="None" value="">
                  กรอกชื่อลูกค้า
                </option>
                {customers.map((item: CustomerInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Checker: </p>
              <Select
                native
                value={checkOut.EmployeeID}
                onChange={handleChange}
                inputProps={{
                  name: "EmployeeID",
                }}
              >
                <option aria-label="None" value="">
                  กรอกชื่อของคุณ
                </option>
                {employees.map((item: EmployeeInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Check-Out Time</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="CheckOutTime"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label="กรุณาเลือกวันที่และเวลา"
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ความเรียบร้อย</p>
              <TextField
                id="Condition"
                variant="outlined"
                type="string"
                size="medium"
                placeholder="มีความเสียหายหรือไม่"
                value={checkOut.Condition || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default CheckOutCreate;