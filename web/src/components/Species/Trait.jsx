import Card, { CardActions, CardContent, CardHeader } from '@material-ui/core/Card';
import Collapse from '@material-ui/core/Collapse';
import IconButton from '@material-ui/core/IconButton';
import { withStyles } from '@material-ui/core/styles';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import classnames from 'classnames';
import PropTypes from 'prop-types';
import React from 'react';
import Variant from './Variant';

const styles = theme => ({
  card: {
    fontFamily: 'Montserrat',
  },
  cardHeaderRoot: { padding: 4 },
  cardHeaderTitle: { fontSize: 14 },
  cardActionRoot: {
    marginTop: -37,
    display: 'flex',
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
        <CardHeader
          title={name}
          classes={{
                  root: classes.cardHeaderRoot,
                  title: classes.cardHeaderTitle,
                }}
        />
        <CardActions classes={{ root: classes.cardActionRoot }} disableActionSpacing>
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
