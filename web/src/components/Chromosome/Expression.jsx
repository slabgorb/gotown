import React from 'react';
import PropTypes from 'prop-types';
import Table, { TableBody, TableCell, TableRow } from 'material-ui/Table';

const _ = require('underscore');

const Expression = ({ expressionMap }) =>
  (
    <Table>
      <TableBody>
        {_.map(expressionMap, (v, k) => (
          <TableRow key={k}>
            <TableCell>{k}</TableCell>
            <TableCell>{v}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );


Expression.propTypes = {
  expressionMap: PropTypes.object.isRequired,
};

export default Expression;
