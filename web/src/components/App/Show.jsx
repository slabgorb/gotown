import { MuiThemeProvider, createMuiTheme, withStyles } from 'material-ui/styles';
import Reboot from 'material-ui/Reboot';
import React from 'react';
import PropTypes from 'prop-types';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import IconButton from 'material-ui/IconButton';
import MenuIcon from 'material-ui-icons/Menu';
import Typography from 'material-ui/Typography';
import Menu, { MenuItem } from 'material-ui/Menu';
import { withRouter } from 'react-router-dom';
import axios from 'axios';


const ralewayFF = ({ fontFamily: 'Raleway' });
const montserratFF = ({ fontFamily: 'Montserrat' });

const styles = () => ({
  flex: { flex: 1 },
  main: { marginTop: '75px' },
  appBar: {},
});

const theme = createMuiTheme({
  palette: {
    primary: {
      light: '#757ce8',
      main: '#3f50b5',
      dark: '#002884',
      contrastText: '#fff',
    },
    secondary: {
      light: '#ff7961',
      main: '#f44336',
      dark: '#ba000d',
      contrastText: '#000',
    },


  },
  typography: {
    fontFamily: 'Raleway',
    body1: ralewayFF,
    body2: ralewayFF,
    display1: ralewayFF,
    display2: ralewayFF,
    display3: ralewayFF,
    display4: ralewayFF,
    caption: ralewayFF,
    headline: montserratFF,
    title: montserratFF,
  },
});


const seed = () => axios.put('/api/seed');

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      anchorEl: null,
    };
    this.handleMenu = this.handleMenu.bind(this);
    this.handleClose = this.handleClose.bind(this);
  }

  handleMenu(event) {
    this.setState({ anchorEl: event.currentTarget });
  }

  handleMenuItem(value) {
    const { history } = this.props;
    history.push(`/${value}`);
    this.handleClose();
  }

  handleClose() {
    this.setState({ anchorEl: null });
  }

  render() {
    const { children, classes } = this.props;
    const { anchorEl } = this.state;
    const open = Boolean(anchorEl);
    return (
      <div>
        <Reboot>
          <MuiThemeProvider theme={theme}>
            <div>
              <AppBar position="fixed" className={classes.appBar}>
                <Toolbar>
                  <Typography variant="title" color="inherit" className={classes.flex}>
                    <a className="logo" href="/">Gotown</a>
                  </Typography>
                  <IconButton onClick={this.handleMenu}>
                    <MenuIcon />
                  </IconButton>
                  <Menu
                    id="menu-bar"
                    open={open}
                    onClose={this.handleClose}
                    anchorEl={anchorEl}
                    anchorOrigin={{
                      vertical: 'top',
                      horizontal: 'right',
                    }}
                  >
                    <MenuItem onClick={() => this.handleMenuItem('words')}>Words</MenuItem>
                    <MenuItem onClick={() => this.handleMenuItem('species')}>Species</MenuItem>
                    <MenuItem onClick={() => this.handleMenuItem('cultures')}>Cultures</MenuItem>
                    <MenuItem onClick={() => this.handleMenuItem('towns')}>Towns</MenuItem>
                    <MenuItem onClick={() => seed()}>Seed</MenuItem>
                  </Menu>
                </Toolbar>
              </AppBar>
              <div className={classes.main}>
                {children}
              </div>
            </div>
          </MuiThemeProvider>
        </Reboot>
      </div>
    );
  }

}


App.propTypes = {
  children: PropTypes.node.isRequired,
  classes: PropTypes.object.isRequired,
  history: PropTypes.object.isRequired,
};

export default withRouter(withStyles(styles)(App));
