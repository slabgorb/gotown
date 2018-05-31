import Table from 'material-ui/Table';
import TableCell from 'material-ui/Table/TableCell';
import TableHead from 'material-ui/Table/TableHead';
import TablePagination from 'material-ui/Table/TablePagination';
import TableRow from 'material-ui/Table/TableRow';
import { withStyles } from 'material-ui/styles';
import PropTypes from 'prop-types';
import React from 'react';

const styles = () => ({
  root: {},
});

class List extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      page: 0,
      rowsPerPage: 25,
    };
    this.handleChangePage = this.handleChangePage.bind(this);
    this.handleChangeRowsPerPage = this.handleChangeRowsPerPage.bind(this);
  }

  handleChangePage(event, page) {
    this.setState({ page });
  }
  handleChangeRowsPerPage(event, rowsPerPage) {
    this.setState({ rowsPerPage });
  }

  render() {
    const { beings } = this.props;
    const { page, rowsPerPage } = this.state;
    const emptyRows = rowsPerPage - Math.min(rowsPerPage, (beings.length - page) * rowsPerPage);
    const table = beings.slice(page * rowsPerPage, (page * rowsPerPage) + rowsPerPage).map(b => (
      <TableRow key={b.id}>
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
        {emptyRows > 0 && (
          <TableRow style={{ height: 48 * emptyRows }}>
            <TableCell colSpan={3} />
          </TableRow>
        )}
        <TablePagination
          colSpan={3}
          count={beings.length}
          rowsPerPage={rowsPerPage}
          rowsPerPageOptions={[25, 50, 75]}
          page={page}
          onChangePage={this.handleChangePage}
          onChangeRowsPerPage={this.handleChangeRowsPerPage}
        />
      </Table>
    );
  }
}


List.propTypes = {
  beings: PropTypes.array.isRequired,
};

export default withStyles(styles)(List);
