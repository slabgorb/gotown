import React from 'react';
import Paper from 'material-ui/Paper';
import { withStyles } from 'material-ui/styles';
import Typography from 'material-ui/Typography';
import PropTypes from 'prop-types';
import inflection from 'inflection';
import cultureApi from './api';
import NameList from '../NameList';

const styles = theme => ({
  root: {
    backgroundColor: theme.palette.background.paper,
    fontFamily: 'Montserrat',
  },
  headline: {
    fontFamily: 'Montserrat',
    marginLeft: '20',
  },
});

class Culture extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.match.params.name,
      names: { family: [], genderNames: {} },
      loaded: false,
    };
    this.get = this.get.bind(this);
  }

  componentDidMount() {
    this.get(this.props.match.params);
  }

  componentWillReceiveProps(nextProps) {
    if (this.props.match.params.name !== nextProps.match.params.name) {
      this.get(nextProps.match.params);
    }
  }

  get({ name }) {
    cultureApi.get(name)
      .then((s) => {
        this.setState({
          name: s.name,
          names: { family: s.family_names, genderNames: s.gender_names },
          loaded: true,
        });
      });
  }


  render() {
    const { classes } = this.props;
    if (!this.state.loaded) {
      return (<div>loading</div>);
    }
    return (
      <div>
        <Paper elevation={4} className={classes.root}>
          <Typography variant="headline" component="h1" className={classes.headline}>
            {inflection.titleize(this.state.name)}
          </Typography>
          <div className="flex-container">
            { (this.state.names.family.length > 0) ? (<NameList title="family names" listItems={this.state.names.family} />) : null}
            {this.state.names.genderNames.map(gn => (<NameList title={inflection.titleize(`${gn.gender} Names`)} listItems={gn.given_names} />))}
          </div>
        </Paper>
      </div>
    );
  }
}

Culture.propTypes = {
  match: PropTypes.object.isRequired,
  classes: PropTypes.object.isRequired,
};

module.exports = withStyles(styles)(Culture);
