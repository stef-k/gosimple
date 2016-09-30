function websocketReducer(state='There are no incomming messages', action) {
  if (action.type === 'INCOMING_MESSGAGE') {
    return Object.assign({}, state, {
      text: action.text
    });
  } else {
    // if the action type is unknown return the state as is
    return state;
  }
};

export default websocketReducer;
