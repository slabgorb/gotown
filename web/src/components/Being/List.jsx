import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';
import { withStyles } from '@material-ui/core/styles';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';

const styles = () => ({
  root: {},
});

const sortString = t => (a, b) => {
  const aS = a[t].toUpperCase();
  const bS = b[t].toUpperCase();
  if (aS < bS) { return -1; }
  if (aS > bS) { return 1; }
  return 0;
};

const sorter = (t) => {
  if (t === 'age') { return (a, b) => a.age - b.age; }
  return sortString(t);
};

class List extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      page: 0,
      rowsPerPage: 25,
      orderBy: 'name',
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
    const { page, rowsPerPage, orderBy } = this.state;
    beings.sort(sorter(orderBy));
    const emptyRows = rowsPerPage - Math.min(rowsPerPage, (beings.length - page) * rowsPerPage);
    const table = beings.slice(page * rowsPerPage, (page * rowsPerPage) + rowsPerPage).map(b => (
      <TableRow key={b.id}>
        <TableCell>{b.name}</TableCell>
        <TableCell>{inflection.titleize(b.gender)}</TableCell>
        <TableCell>{b.age}</TableCell>
        <TableCell>{inflection.titleize(b.species)}</TableCell>
        <TableCell>{inflection.titleize(b.culture)}</TableCell>
      </TableRow>
    ));
    return (
      <div>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell>Gender</TableCell>
              <TableCell>Age</TableCell>
              <TableCell>Species</TableCell>
              <TableCell>Culture</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {table}
            {emptyRows > 0 && (
              <TableRow style={{ height: 48 * emptyRows }}>
                <TableCell colSpan={5} />
              </TableRow>
            )}
          </TableBody>

        </Table>
        <TablePagination
          component="div"
          colSpan={3}
          count={beings.length}
          rowsPerPage={rowsPerPage}
          rowsPerPageOptions={[25, 50, 75]}
          page={page}
          onChangePage={this.handleChangePage}
          onChangeRowsPerPage={this.handleChangeRowsPerPage}
        />
      </div>
    );
  }
}


List.propTypes = {
  beings: PropTypes.array.isRequired,
};

export default withStyles(styles)(List);
