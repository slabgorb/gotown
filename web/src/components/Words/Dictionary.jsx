import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import { Card, CardContent, CardHeader, CardActions, IconButton, Collapse, Typography } from '@material-ui/core';
import ExpandMore from '@material-ui/icons/ExpandMore';
import NameList from '../NameList';
import classnames from 'classnames';
import inflection from 'inflection';

const styles = theme => ({
  cardContent: {
    backgroundColor: theme.palette.background.paper,
  },
  titleRoot: {
    display: 'flex',
  },
  list: {
  },
  listItem: {
  },
  cardActionRoot: {
    display: 'flex',
  },
  root: {
    marginBottom: theme.spacing.unit * 3,  
  }
});


class Dictionary extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      expanded: false,
    }
    this.handleExpandClick = this.handleExpandClick.bind(this);
  }

  handleExpandClick() {
    this.setState({ expanded: !this.state.expanded });
  }

  render() {
    const { k, dictionary, classes } = this.props;
    const { expanded } = this.state;
    const title = (
      <div className={classes.titleRoot}>
        <IconButton
          className={classnames(classes.expand, {
            [classes.expandOpen]: expanded,
          })}
          onClick={this.handleExpandClick}
          aria-expanded={this.state.expanded}
          aria-label="Show more"
        >
          <ExpandMore />
        </IconButton>
        <Typography variant="display1" key={k}>{inflection.titleize(k)}</Typography>
      </div>
    )
    return (<Card className={classes.root}>
      <CardHeader title={title} />
      <Collapse in={expanded} timeout="auto" unmountOnExit>
        <CardContent>
          <NameList
            key={k}
            listItems={dictionary}
          />
        </CardContent>
      </Collapse>
    </Card>)
  }
}

Dictionary.propTypes = {
  k: PropTypes.string.isRequired,
  dictionary: PropTypes.array.isRequired,
  classes: PropTypes.object.isRequired,
}

module.exports = withStyles(styles)(Dictionary);