import Table from '@material-ui/core/Table';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import PropTypes from 'prop-types';
import React from 'react';

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
