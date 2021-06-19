import React from "react";
import Form from "react-bootstrap/Form";
import Col from "react-bootstrap/Col";
import Row from "react-bootstrap/Row";
import Button from "react-bootstrap/Button";
import Container from "react-bootstrap/Container";
import Toast from "./Toaster";
import axios from "axios";

type AdminState = {
  validated: boolean;
  id: string;
  buttonDisabled: boolean;
  toastShow: boolean;
  toastText: string;
};

class Admin extends React.Component<{}, any> {
  constructor(props: {}) {
    super(props);
    this.state = {
      validated: false,
      id: "",
      buttonDisabled: false,
      toastShow: false,
      toastText: "asd",
    };

    this.validated = this.validated.bind(this);
    this.changeHandler = this.changeHandler.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  validated(): boolean {
    console.log("validated", this.state.validated);
    return this.state.validated;
  }


  async handleSubmit(event: React.ChangeEvent<any>) {
    event.preventDefault();
    if(event.currentTarget.checkValidity() === false) {
        event.stopPropagation();
        return;
    };

    this.setState({buttonDisabled: true, toastShow: false})

    // create the data
    const data = new FormData()
    data.append('file', this.state.file);

    try {
      const res = await axios({
        url: `http://localhost:8080/images/${this.state.id}`,
        method: 'POST',
        data: data,
        headers: {
          'Content-Type': 'multipart/form-data',
        }
      })
      
      const toastText = res.status === 200 ? "Uploaded file" : "Unable to upload file. Error:" +res.statusText;
      this.setState({buttonDisabled: false, toastShow: true, toastText: toastText});
    } catch(error) {
      console.log("Err" + error);
      this.setState({buttonDisabled: false, toastShow: true, toastText: "Unable to upload file. " + error});
    }
  }

  changeHandler(event: any) {
    if (event.target.name === "file") {
      this.setState({
        file: event.target.files[0],
        toastShow: false,
      });
      return;
    }

    this.setState({
      [event.target.name]: event.target.value,
      toastShow: false,
    });
  }

  render() {
    return (
      <div>
        <h1 style={{marginBottom: "40px"}}>Admin</h1>
        <Container className="text-left">
          <Form noValidate validated={this.validated()} onSubmit={this.handleSubmit}>
            <Form.Group as={Row} controlId="productID">
              <Form.Label column sm="2">Product ID:</Form.Label>
              <Col sm="6">
                <Form.Control type="text" name="id" placeholder="" required style={{width: "80px"}} value={this.state.id} onChange={this.changeHandler}/>
                <Form.Text className="text-muted">Enter the product id to upload an image for</Form.Text>
                <Form.Control.Feedback type="invalid">Please provide a product ID.</Form.Control.Feedback>
              </Col>
              <Col sm="4">
                <Toast state={this.state.toastShow} message={this.state.toastText}/>
              </Col>
            </Form.Group>
            <Form.Group as={Row}>
              <Form.Label column sm="2">File:</Form.Label>
              <Col sm="10">
                <Form.Control type="file" name="file" placeholder="" required onChange={this.changeHandler}/>
                <Form.Text className="text-muted">Image to associate with the product</Form.Text>
                <Form.Control.Feedback type="invalid">Please select a file to upload.</Form.Control.Feedback>
              </Col>
            </Form.Group>
            <Button  type="submit" disabled={this.state.buttonDisabled}>Submit form</Button>
          </Form>
        </Container>
      </div>
    );
  }
}


export default Admin;