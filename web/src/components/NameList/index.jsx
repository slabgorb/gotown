import Card, { CardContent, CardHeader } from '@material-ui/core/Card';
import List, { ListItem } from '@material-ui/core/List';
import { withStyles } from '@material-ui/core/styles';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';

const _ = require('underscore');

const styles = theme => ({
  root: {
  },
  cardContent: {
    backgroundColor: theme.palette.background.paper,
  },
  cardHeader: {
    fontFamily: 'Raleway',
    fontSize: '14',
    subheader: {
      fontFamily: 'Raleway',
    },
  },
  list: {
  },
  listItem: {
    fontFamily: 'Raleway',
    fontSize: '12',
  },
});

const NameList = ({
  classes,
  title,
  listItems,
}) => (
  <Card className={classes.root}>
    <CardHeader
      title={inflection.titleize(title)}
      className={classes.cardHeader}
    />
    <CardContent className={classes.cardContent}>
      <List className={classes.list} dense disablePadding>
        { _.uniq(listItems).map(f =>
          (<ListItem className={classes.listItem} key={f}>{f}</ListItem>))}
      </List>
    </CardContent>
  </Card>
);

NameList.propTypes = {
  title: PropTypes.string.isRequired,
  listItems: PropTypes.array.isRequired,
  classes: PropTypes.object.isRequired,
};
NameList.defaultProps = {
};


module.exports = withStyles(styles)(NameList);
