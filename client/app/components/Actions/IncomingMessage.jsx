// an example action using an action creator function
// to dispatch this, use store.dispatch(exampleAction('the text'));

export default function incomingMessage(text) {
  return {
    type: 'INCOMING_MESSGAGE',
    text: text
  };
}

