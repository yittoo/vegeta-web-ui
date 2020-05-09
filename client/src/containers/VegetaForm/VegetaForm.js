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
          value: undefined,
          warnRequired: false,
        },
        target: {
          name: "target",
          value: undefined,
          warnRequired: false,
        },
        method: {
          name: "method",
          value: undefined,
          warnRequired: false,
        },
        duration: {
          name: "duration",
          value: undefined,
          warnRequired: false,
        },
        frequency: {
          name: "freq",
          value: undefined,
          warnRequired: false,
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
          warnRequired: false,
          value,
        },
      },
    });
  };

  validateFields = () => {
    const { formData } = this.state;
    let isValid = true;

    if (!formData.attackName.value) {
      isValid = false;
      formData.attackName.warnRequired = true;
    }
    if (!formData.target.value || !formData.target.value.startsWith("http")) {
      isValid = false;
      formData.target.warnRequired = true;
    }
    if (!formData.method.value) {
      isValid = false;
      formData.method.warnRequired = true;
    }
    if (!formData.duration.value || +formData.duration.value < 1) {
      isValid = false;
      formData.duration.warnRequired = true;
    }
    if (!formData.frequency.value || +formData.frequency.value < 1) {
      isValid = false;
      formData.frequency.warnRequired = true;
    }

    return isValid;
  };

  onSubmit = () => {
    const { formData } = this.state;
    let body = {};

    const isValid = this.validateFields();
    if (!isValid) return;

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
          <h3 className={s.NewTestTitle}>New Test</h3>
          <Col span={24}>
            <Input
              value={attackName.value}
              onChange={(e) =>
                this.onChangeHandler("attackName", e.target.value)
              }
              placeholder="Name of test, ex: Test MySite.com"
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
