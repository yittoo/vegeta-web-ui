import React, { useReducer, createContext } from "react";

const ViewResultsContext = createContext();

ViewResultsContext.displayName = "ViewResultsAs";

const initialState = {
  options: Object.freeze({
    json: "json",
    graph: "graph",
  }),
  value: "json",
};

const actionTypes = {
  SET_VIEW_RESULTS_AS: "ViewResultsAs.SET_VIEW_RESULTS_AS",
  TOGGLE_VIEW_RESULTS_AS: "ViewResultsAs.TOGGLE_VIEW_RESULTS_AS",
};

const reducer = (state, { type, payload }) => {
  switch (type) {
    case actionTypes.SET_VIEW_RESULTS_AS:
      return {
        ...state,
        value: payload,
      };
    case actionTypes.TOGGLE_VIEW_RESULTS_AS:
      return {
        ...state,
        value:
          state.value === initialState.options.json
            ? initialState.options.graph
            : initialState.options.json,
      };
    default:
      return state;
  }
};

const actions = {
  setViewResultsAs: (to) => {
    return {
      type: actionTypes.SET_VIEW_RESULTS_AS,
      payload: to,
    };
  },
  toggleViewResultsAs: () => {
    return {
      type: actionTypes.TOGGLE_VIEW_RESULTS_AS,
    };
  },
};

export const ViewResultsConsumer = ViewResultsContext.Consumer;

export const ViewResultsProvider = ({ children }) => {
  const [state, dispatch] = useReducer(reducer, initialState);
  return (
    <ViewResultsContext.Provider value={{ state, dispatch, actions }}>
      {children}
    </ViewResultsContext.Provider>
  );
};
