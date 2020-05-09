import React from "react";
import { Row, Col, Tabs } from "antd";
import moment from "moment";

import s from "./Results.module.scss";
import { ViewResultsConsumer } from "../../context";
import { ResultsTable } from "../../components";

const { TabPane } = Tabs;

export class Results extends React.Component {
  renderTabs = (data) => {
    if (!data || data.length === 0) {
      return (
        <Tabs>
          <TabPane key="no_attack" tab="Test name">
            Results will go here
          </TabPane>
        </Tabs>
      );
    }

    return (
      <Tabs>
        {data.map((a) => (
          <TabPane key={a.timeStampMilisecond} tab={a.name}>
            {this.renderResAsHtml(a.html)}
          </TabPane>
        ))}
      </Tabs>
    );
  };

  renderResAsHtml = (html) => {
    return <iframe className={s.Graph} srcDoc={html} />;
  };

  renderResAsTable = (data) => {
    if (!data || data.length === 0) {
      return <ResultsTable />;
    }
    const tableData = data.map((a) => {
      const { json } = a;
      const { latencies } = json;
      return {
        key: a.timeStampMilisecond,
        name: a.name,
        date: moment(a.timeStampMilisecond).format("HH:mm:ss - DD.MM.YYYY"),
        duration: (json.duration / 1000000000).toFixed(2),
        rate: json.rate.toFixed(1),
        latency_mean: (latencies.mean / 1000000000).toFixed(2) + " s",
        latency_50: (latencies["50th"] / 1000000000).toFixed(2) + " s",
        latency_95: (latencies["95th"] / 1000000000).toFixed(2) + " s",
        latency_max: (latencies.max / 1000000000).toFixed(2) + " s",
        requests_total: json.requests,
        requests_success: json.status_codes["200"],
      };
    });
    return (
      <div>
        <h3>Tabular Results</h3>
        <ResultsTable data={tableData} />
      </div>
    );
  };

  render() {
    const { data } = this.props;
    return (
      <section className={s.Results}>
        <ViewResultsConsumer>
          {({ state }) => {
            if (state.value === state.options.graph) {
              return this.renderTabs(data);
            } else {
              return this.renderResAsTable(data);
            }
          }}
        </ViewResultsConsumer>
      </section>
    );
  }
}
