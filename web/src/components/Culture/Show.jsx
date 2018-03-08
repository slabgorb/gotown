import React from 'react';
import Paper from 'material-ui/Paper';
import Card, { CardHeader, CardContent } from 'material-ui/Card';
import Typography from 'material-ui/Typography';
import PropTypes from 'prop-types';
import cultureApi from './api';

class Culture extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: props.match.params.name,
      names: { family: [], genderNames: {} },
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
        });
      });
  }

  render() {
    return (
      <div>
        <Paper elevation={4}>
          <Typography variant="headline" component="h1">
            {this.state.name}
          </Typography>
          <Card>
            <CardHeader title="Family Names" />
            <CardContent>
              { this.state.names.family.map(f => (<p>{f}</p>))}
            </CardContent>
          </Card>
        </Paper>
      </div>
    );
  }
}

Culture.propTypes = {
  match: PropTypes.object.isRequired,
};

module.exports = Culture;
