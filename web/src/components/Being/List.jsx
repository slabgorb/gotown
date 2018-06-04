import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';
import TableSortLabel from '@material-ui/core/TableSortLabel';
import { withStyles } from '@material-ui/core/styles';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';

const styles = () => ({
  root: {},
});

const sortString = (t, dir) => (a, b) => {
  let up = -1;
  let down = 1;
  if (dir === 'desc') {
    up = 1;
    down = -1;
  }
  const aS = a[t].toUpperCase();
  const bS = b[t].toUpperCase();
  if (aS < bS) { return up; }
  if (aS > bS) { return down; }
  return 0;
};

const sorter = (t, dir) => {
  if (t === 'age') { return dir === 'asc' ? ((a, b) => a.age - b.age) : ((a, b) => b.age - a.age); }
  return sortString(t, dir);
};

class List extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      page: 0,
      rowsPerPage: 25,
      orderBy: 'name',
      dir: 'asc',
    };
    this.handleChangePage = this.handleChangePage.bind(this);
    this.handleChangeRowsPerPage = this.handleChangeRowsPerPage.bind(this);
    this.header = this.header.bind(this);
    this.headers = this.headers.bind(this);
    this.handleChangeSortFactory = this.handleChangeSortFactory.bind(this);
  }

  handleChangePage(event, page) {
    this.setState({ page });
  }
  handleChangeRowsPerPage(event, rowsPerPage) {
    this.setState({ rowsPerPage });
  }

  headers() {
    const headerEntries = ['name', 'gender', 'age', 'species', 'culture'];
    return (
      <TableRow>
        {headerEntries.map(he => this.header(he))}
      </TableRow>
    );
  }

  handleChangeSortFactory(orderBy) {
    return () => {
      let dir = 'desc';
      if (this.state.orderBy === orderBy && this.state.dir === 'desc') {
        dir = 'asc';
      }
      this.setState({
        orderBy,
        dir,
      });
    };
  }

  header(label) {
    const { orderBy, dir } = this.state;
    return (
      <TableCell sortDirection={orderBy === label ? dir : false}>
        <TableSortLabel
          active={orderBy === label}
          direction={dir}
          onClick={this.handleChangeSortFactory(label)}
        >
          {inflection.titleize(label)}
        </TableSortLabel>
      </TableCell>
    );
  }

  render() {
    const { beings } = this.props;
    const { page, rowsPerPage, orderBy, dir } = this.state;
    beings.sort(sorter(orderBy, dir));
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
            {this.headers()}
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
