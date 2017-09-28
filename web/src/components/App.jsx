import React from 'react';
import Species from 'components/Species.jsx';
import { Area } from 'components/Area.jsx'
import Being from 'components/Being.jsx'
import Nav from 'components/Nav.jsx';
import { Route } from 'react-router-dom'

const App = (props) => {
  return   (
      <div>
        <Species name="human"/>
      </div>
    )
}


//
// class App extends React.Component {
//   render() {
//     return(
//     <Area/>
//     )
//   }
// }

module.exports = App
