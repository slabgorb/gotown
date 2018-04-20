import React from 'react';
import PropTypes from 'prop-types';
import ExpandMoreIcon from 'material-ui-icons/ExpandMore';
import ExpansionPanel, {
  ExpansionPanelSummary,
  ExpansionPanelDetails,
} from 'material-ui/ExpansionPanel';
import List, { ListItem, ListItemText } from 'material-ui/List';
import Typography from 'material-ui/Typography';
import inflection from 'inflection';

const MenuList = ({
  classes,
  heading,
  list,
  handleClick,
}) => (
  <ExpansionPanel>
    <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
      <Typography className={classes.heading}>{heading}</Typography>
    </ExpansionPanelSummary>
    <ExpansionPanelDetails>
      <List component="nav">
        {list.map(item => (
          <ListItem button divider key={item} onClick={() => handleClick(item)}>
            <ListItemText primary={inflection.titleize(item)} />
          </ListItem>
          ))}
      </List>
    </ExpansionPanelDetails>
  </ExpansionPanel>
);

MenuList.propTypes = {
  classes: PropTypes.object.isRequired,
  heading: PropTypes.node.isRequired,
  list: PropTypes.array.isRequired,
  handleClick: PropTypes.func.isRequired,
}

module.exports = MenuList;