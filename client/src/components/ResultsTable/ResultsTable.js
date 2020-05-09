import React from "react";
import { Table } from "antd";
import { SettingTwoTone } from "@ant-design/icons";

import s from "./ResultsTable.module.scss";

const columns = [
  {
    title: "Test Name",
    dataIndex: "name",
    key: "name",
  },
  {
    title: "Date",
    dataIndex: "date",
    key: "date",
  },
  {
    title: "Duration",
    dataIndex: "duration",
    key: "duration",
  },
  {
    title: "Requests per second",
    dataIndex: "rate",
    key: "rate",
  },
  {
    title: "Mean Latency",
    dataIndex: "latency_mean",
    key: "latency_mean",
  },
  {
    title: "Latency (50%)",
    dataIndex: "latency_50",
    key: "latency_50",
  },
  {
    title: "Latency (95%)",
    dataIndex: "latency_95",
    key: "latency_95",
  },
  {
    title: "Maximum Latency",
    dataIndex: "latency_max",
    key: "latency_max",
  },
  {
    title: "Total requests",
    dataIndex: "requests_total",
    key: "requests_total",
  },
  {
    title: "Successful requests",
    dataIndex: "requests_success",
    key: "requests_success",
  },
];

export class ResultsTable extends React.Component {
  render() {
    return (
      <Table
        className={s.Table}
        columns={columns}
        dataSource={this.props.data}
        pagination={false}
        locale={{
          emptyText: (
            <div className={s.EmptyResults}>
              <p className={s.EmptyResults__Text}>Your results will go here</p>
              <SettingTwoTone
                twoToneColor="#a901db"
                className={s.EmptyResults__Icon}
              />
            </div>
          ),
        }}
      />
    );
  }
}
