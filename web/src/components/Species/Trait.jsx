import React from 'react';
import Card, { CardHeader, CardContent, CardActions } from 'material-ui/Card';
import PropTypes from 'prop-types';
import Collapse from 'material-ui/transitions/Collapse';
import ExpandMoreIcon from 'material-ui-icons/ExpandMore';
import { withStyles } from 'material-ui/styles';
import IconButton from 'material-ui/IconButton';
import classnames from 'classnames';
import Variant from './Variant';

const styles = theme => ({
  card: {
    minWidth: 275,
    fontFamily: 'Montserrat',
  },
  title: {
    fontSize: 12,
  },

  expand: {
    transform: 'rotate(0deg)',
    transition: theme.transitions.create('transform', {
      duration: theme.transitions.duration.shortest,
    }),
    marginLeft: 'auto',
  },
  expandOpen: {
    transform: 'rotate(180deg)',
  },
  actions: {
    display: 'flex',
  },
});

class Trait extends React.Component {
  constructor(props) {
    super(props);
    this.state = { expanded: false };
    this.handleExpandClick = this.handleExpandClick.bind(this);
  }

  handleExpandClick() {
    this.setState({ expanded: !this.state.expanded });
  }

  render() {
    const { name, variants, classes } = this.props;
    return (
      <Card className={classes.card}>
        <CardHeader title={name} className={classes.title} />
        <CardActions className={classes.actions} disableActionSpacing>
          <IconButton
            className={classnames(classes.expand, {
                [classes.expandOpen]: this.state.expanded,
              })}
            onClick={this.handleExpandClick}
            aria-expanded={this.state.expanded}
            aria-label="Show more"
          >
            <ExpandMoreIcon />
          </IconButton>
        </CardActions>
        <Collapse in={this.state.expanded} timeout="auto" unmountOnExit>
          <CardContent>
            {
              variants.map(v =>
                (<Variant name={v.name} match={v.match} key={v.name + v.match} />))
            }
          </CardContent>
        </Collapse>
      </Card>
    );
  }
}

Trait.propTypes = {
  name: PropTypes.string.isRequired,
  variants: PropTypes.array.isRequired,
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Trait);
