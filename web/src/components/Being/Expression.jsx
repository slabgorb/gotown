import React from 'react';

const Expression = props =>
  (
    <div className="being-expression">
      {_.map(props.expression, expressionMap)}
    </div>
  );
