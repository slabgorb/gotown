import React from 'react';
import Card, { CardContent } from 'material-ui/Card';
import Expression from '../Chromosome';

const _ = require('underscore');

const Being = props =>
  (
    <div>
      <Card className="being">
        <CardContent>
          <div className="being-name">{props.being.name.display_name}</div>
          <div className="being-age">{props.being.age}</div>
          <div className="being-gender">{props.being.gender}</div>
          <Expression expression={props.being.expression} />
          <Chromosome chromosome={props.being.chromosome} />
        </CardContent>
      </Card>
      <br />
    </div>
  );
const expressionMap = (v, k) =>
  (
    <div key={k} className="key-value">
      <div>{k}</div>
      <div>{v}</div>
    </div>
  );


const geneMap = gene => <div key={gene} className="being-chromosome-gene" style={{ backgroundColor: `#${gene}` }} />;
const Chromosome = props =>
  (
    <div className="being-chromosome">
      {props.chromosome.genes.map(geneMap)}
    </div>

  );

module.exports = Being;
