import React from "react";
import Table from "react-bootstrap/Table";
import axios from "axios";

interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
  sku: string;
  created_at: string;
  updated_at: string;
}

class CoffeList extends React.Component<{}, { products: Product[] }> {
  constructor(props: {}) {
    super(props);
    this.getData();
    this.state = { products: [] };
    this.getData = this.getData.bind(this);
  }

  async getData() {
    try {
      const res = await axios.get<Product[]>('http://localhost:1323/products/');
      console.log(res);
      this.setState({ products: res.data });
    } catch (error) {
      console.error(error);
    }
  }

  setProducts() {
    let table = [];
    for (let i = 0; i < this.state.products.length; i++) {
      table.push(
        <tr key={i}>
          <td>{this.state.products[i].name}</td>
          <td>{this.state.products[i].price}</td>
          <td>{this.state.products[i].sku}</td>
        </tr>
      );
    }

    return table;
  }

  render() {
    return (
      <div>
        <h1 style={{ marginBottom: "40px" }}>Menu</h1>
        <Table>
          <thead>
            <tr>
              <th>Name</th>
              <th>Price</th>
              <th>SKU</th>
            </tr>
          </thead>
          <tbody>{this.setProducts()}</tbody>
        </Table>
      </div>
    );
  }
}

export default CoffeList;