import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import { Table, TableRow, TableCell } from 'material-ui/Table';

const styles = () => ({
  root: {},
});

const List = ({ beings }) => {
  const table = beings.map(b => (
    <TableRow>
      <TableCell>{b.name}</TableCell>
      <TableCell>{b.gender}</TableCell>
      <TableCell>{b.age}</TableCell>
    </TableRow>

  ));
  return (
    <Table>
      {table}
    </Table>

  );
};

List.propTypes = {
  beings: PropTypes.array.isRequired,
};

export default withStyles(styles)(List);
