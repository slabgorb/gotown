import React from 'react';
import Species from './Species';
import { Area } from './Area.jsx'
// import Being from 'components/Being.jsx'
// import Nav from 'components/Nav.jsx';
//import { Route } from 'react-router-dom'

const App = () =>
  (
    <div>
      <Area/>
      <Species name="human" />
    </div>
  );



//
// class App extends React.Component {
//   render() {
//     return(
//     <Area/>
//     )
//   }
// }

module.exports = App
