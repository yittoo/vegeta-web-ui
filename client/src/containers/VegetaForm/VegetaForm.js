import React, { Component } from "react";
import { Row, Col, Select, Tabs, Input, Button } from "antd";

import s from "./VegetaForm.module.scss";

export class VegetaForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      formData: {
        attackName: {
          name: "attackName",
          value: "attack",
        },
        target: {
          name: "target",
          value: "https://sozerdesign.com",
        },
        method: {
          name: "method",
          value: "GET",
        },
        duration: {
          name: "duration",
          value: "1",
        },
        frequency: {
          name: "freq",
          value: "5",
        },
      },
      pendingResult: false,
      loaderData: {
        attackLength: 10,
        defaultTimeout: 10,
      },
    };
  }

  onChangeHandler = (key, value) => {
    this.setState({
      ...this.state,
      formData: {
        ...this.state.formData,
        [key]: {
          ...this.state.formData[key],
          value,
        },
      },
    });
  };

  onSubmit = () => {
    const { formData } = this.state;
    let body = {};

    Object.keys(formData).forEach((key) => {
      body[formData[key].name] = formData[key].value;
    });

    fetch("http://localhost:8000/vegeta", {
      method: "POST",
      body: JSON.stringify(body),
    }).then(async (res) => {
      const json = await res.json();
      this.props.onAddNewResult({
        name: this.state.formData.attackName.value || "Boom",
        payload: {
          html: json.AsGraphHTML,
          json: JSON.parse(json.AsGraphJSON),
          timeStampSecond: json.TimeOfAttack,
        },
      });
    });
  };

  render() {
    const {
      attackName,
      target,
      method,
      frequency,
      duration,
    } = this.state.formData;
    return (
      <section className={s.Section}>
        <Row align="stretch" justify="center" className={s.Form}>
          <Col span={24}>
            <Input
              value={attackName.value}
              onChange={(e) =>
                this.onChangeHandler("attackName", e.target.value)
              }
              placeholder="Name of attack, default: Boom"
            />
          </Col>
          <Col span={24}>
            <Input
              value={target.value}
              onChange={(e) => this.onChangeHandler("target", e.target.value)}
              placeholder="Target URL ex: https://google.com"
            />
          </Col>
          <Col span={24}>
            <Select
              value={method.value}
              onChange={(value) => this.onChangeHandler("method", value)}
              className={s.Select}
              placeholder="HTTP Method"
            >
              <Select.Option value={"GET"}>GET</Select.Option>
              <Select.Option value={"POST"}>POST</Select.Option>
            </Select>
          </Col>
          <Col span={24}>
            <Input
              value={duration.value}
              onChange={(e) => this.onChangeHandler("duration", e.target.value)}
              placeholder="Duration (s)"
            />
          </Col>
          <Col span={24}>
            <Input
              value={frequency.value}
              onChange={(e) =>
                this.onChangeHandler("frequency", e.target.value)
              }
              placeholder="Requests per second"
            />
          </Col>
        </Row>
        <Row>
          <Col span={24}>
            <Button className={s.Submit} onClick={this.onSubmit} type="primary">
              Submit
            </Button>
          </Col>
        </Row>
      </section>
    );
  }
}
