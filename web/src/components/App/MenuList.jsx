import React from 'react';
import PropTypes from 'prop-types';
import ExpandMoreIcon from 'material-ui-icons/ExpandMore';
import ExpansionPanel, {
  ExpansionPanelSummary,
  ExpansionPanelDetails,
} from 'material-ui/ExpansionPanel';
import DeleteForeverIcon from 'material-ui-icons/DeleteForever';
import IconButton from 'material-ui/IconButton';
import List, { ListItem, ListItemText } from 'material-ui/List';
import Typography from 'material-ui/Typography';
import inflection from 'inflection';

const deleteIcon = (classes, handleDelete, item) => (
  <IconButton className={classes.deleteButton} onClick={e => handleDelete(e, item.id)}>
    <DeleteForeverIcon />
  </IconButton>
);

const MenuList = ({
  classes,
  heading,
  list,
  handleClick,
  handleDelete,
  deletable,
}) => (
  <ExpansionPanel>
    <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
      <Typography className={classes.heading}>{heading}</Typography>
    </ExpansionPanelSummary>
    <ExpansionPanelDetails>
      <List component="nav">
        {list.map(item => (
          <ListItem button divider key={item.id} onClick={() => handleClick(item.id)}>
            <ListItemText primary={inflection.titleize(item.name)} />
            {deletable ? deleteIcon(classes, handleDelete, item) : (<div />)}
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
  handleDelete: PropTypes.func,
  deletable: PropTypes.bool,
};

MenuList.defaultProps = {
  handleDelete: () => {},
  deletable: false,
};

module.exports = MenuList;
