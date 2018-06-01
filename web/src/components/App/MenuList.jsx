import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import IconButton from '@material-ui/core/IconButton';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Typography from '@material-ui/core/Typography';
import DeleteForeverIcon from '@material-ui/icons/DeleteForever';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import inflection from 'inflection';
import PropTypes from 'prop-types';
import React from 'react';

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
            {item.icon ? (<img src={item.icon} alt="icon" />) : (<div />)}
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
