import React from 'react';
var _ = require('underscore');

const Being = (props) => {
  (
    <div className="being">
      <div className="being-name">{this.props.being.name.display_name}</div>
      <div className="being-age">{this.props.being.age}</div>
      <Chromosome chromosome={this.props.being.chromosome}/>
    </div>
  )
}

const Chromosome = (props) => {
  var cDisplay = _.map(this.props.chromosome.genes, function(gene) {

    var s = {
      backgroundColor: "#" + gene,
    }
    return (
      <div key={gene} className="being-chromosome-gene" style={s}></div>
    )
  })

  (
    <div className="being-chromosome">
      {cDisplay}
    </div>

  )
}

module.exports = Being
