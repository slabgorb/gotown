import React from 'react';
import PropTypes from 'prop-types';
import List, { ListItem } from 'material-ui/List';
import Card, { CardHeader, CardContent } from 'material-ui/Card';
import { withStyles } from 'material-ui/styles';
import inflection from 'inflection';

const styles = theme => ({
  root: {
    width: '150',
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
  subtitle,
  listItems,
}) => (
  <Card className={classes.root}>
    <CardHeader
      title={inflection.titleize(title)}
      subtitle={inflection.titleize(subtitle)}
      className={classes.cardHeader}
    />
    <CardContent className={classes.cardContent}>
      <List className={classes.list} dense disablePadding>
        { listItems.map(f => (<ListItem className={classes.listItem} key={f}>{f}</ListItem>))}
      </List>
    </CardContent>
  </Card>
);

NameList.propTypes = {
  subtitle: PropTypes.string,
  title: PropTypes.string.isRequired,
  listItems: PropTypes.array.isRequired,
  classes: PropTypes.object.isRequired,
};
NameList.defaultProps = {
  subtitle: '',
};


module.exports = withStyles(styles)(NameList);
