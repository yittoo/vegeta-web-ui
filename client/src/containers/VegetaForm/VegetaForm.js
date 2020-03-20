import React, { Component } from "react";
import { Row, Col, Select, Tabs } from "antd";

import s from "./VegetaForm.module.scss";

export class VegetaForm extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }
  render() {
    return (
      <section>
        <Row>
          <Col xs={12}></Col>
        </Row>
        {/* <Select>
          <Select.Option value={1}>Hi</Select.Option>
        </Select> */}
      </section>
    );
  }
}
