import React, { useState } from "react";
import s from "./Drawer.module.scss";

export const Drawer = (props) => {
  const [beforeDragClientX, setBeforeDragClientX] = useState(0);
  const [beforeDragLeftWidth, setBeforeDragLeftWidth] = useState(
    props.originalFromLeft
  );
  const [currentLeftWidth, setCurrentLeftWidth] = useState(
    props.originalFromLeft
  );

  const onDrag = (e) => {
    if (e.clientX > 0)
      setCurrentLeftWidth(e.clientX - beforeDragClientX + beforeDragLeftWidth);
  };

  const onDragStart = (e) => {
    setBeforeDragLeftWidth(currentLeftWidth);
    setBeforeDragClientX(e.clientX);
  };

  return (
    <div className={s.Drawer}>
      <div
        className={[s.Panel, s.LeftPanel].join(" ")}
        style={{ width: currentLeftWidth }}
      >
        {props.leftChildren}
      </div>
      <div className={s.Splitter} onDrag={onDrag} onDragStart={onDragStart}>
        <span></span>
        <span className={s.Thick}></span>
        <span></span>
      </div>
      <div className={[s.Panel, s.RightPanel].join(" ")}>
        {props.rightChildren}
      </div>
    </div>
  );
};
