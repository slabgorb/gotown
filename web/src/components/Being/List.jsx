import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Table from 'material-ui/Table';
import TableCell from 'material-ui/Table/TableCell';
import TableRow from 'material-ui/Table/TableRow';
import TableHead from 'material-ui/Table/TableHead';

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
      <TableHead>
        <TableRow>
          <TableCell>Name</TableCell>
          <TableCell>Gender</TableCell>
          <TableCell>Age</TableCell>
        </TableRow>

      </TableHead>
      {table}
    </Table>

  );
};

List.propTypes = {
  beings: PropTypes.array.isRequired,
};

export default withStyles(styles)(List);
