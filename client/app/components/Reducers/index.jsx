import Redux, {combineReducers, createStore} from 'redux'
import websocketReducer from './WebsocketReducer';

// the main reducer that combines all other reducers
// here message is the store object where websocketReducer is responsible for
// to get the message text use message.text, for details check in Actions/IncomingMessage
const reducer = combineReducers({
  message: websocketReducer,
});

const initialState = {
  message: "Nothing arrived from websocket yet"
};

// application's store
// uncomment for production
// const store = createStore(reducer);
// development
// comment for production
const store = createStore(reducer, initialState, window.devToolsExtension && window.devToolsExtension());

export default store;
