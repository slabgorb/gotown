import { MuiThemeProvider, createMuiTheme, withStyles } from 'material-ui/styles';
import Reboot from 'material-ui/Reboot';
import React from 'react';
import PropTypes from 'prop-types';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import IconButton from 'material-ui/IconButton';
import Button from 'material-ui/Button';
import MenuIcon from 'material-ui-icons/Menu';
import CloseIcon from 'material-ui-icons/Close';
import Typography from 'material-ui/Typography';
import Drawer from 'material-ui/Drawer';
import { withRouter } from 'react-router-dom';
import axios from 'axios';
import { SpeciesList } from '../Species';
import { CulturesList } from '../Culture';
import { WordsList } from '../Words';
import { NamersList } from '../Namer';

const ralewayFF = ({ fontFamily: 'Raleway' });
const montserratFF = ({ fontFamily: 'Montserrat' });

const styles = () => ({
  flex: { flex: 1 },
  main: { marginTop: '75px' },
  appBar: {},
  button: {
    color: 'white',
  },
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
      drawerOpen: false,
    };
    this.handleDrawer = this.handleDrawer.bind(this);
  }

  handleDrawer(drawerOpen) {
    return () => {
      this.setState({ drawerOpen });
    };
  }

  handleMenuItem(value) {
    const { history } = this.props;
    return () => {
      history.push(`/${value}`);
      this.handleDrawer(false)();
    };
  }


  render() {
    const { children, classes } = this.props;
    const { drawerOpen } = this.state;
    return (
      <div>
        <Reboot>
          <MuiThemeProvider theme={theme}>
 
            <div>
              <Drawer
                open={drawerOpen}
                anchor="right"
                onClose={this.handleDrawer(false)}
              >
                <div>
                  <AppBar position="static" className={classes.appBar}>
                    <Toolbar>
                      <IconButton color="default" onClick={this.handleDrawer(false)}>
                        <CloseIcon className={classes.button} />
                      </IconButton>
                      <Button className={classes.button} onClick={this.handleMenuItem('towns')}>Towns</Button>
                      <Button className={classes.button} onClick={this.handleMenuItem('seed')}>Seed</Button>
                    </Toolbar>
                  </AppBar>
                  <SpeciesList handleClick={v => this.handleMenuItem(v)} />
                  <CulturesList handleClick={v => this.handleMenuItem(v)} />
                  <WordsList handleClick={v => this.handleMenuItem(v)} />
                  <NamersList handleClick={v => this.handleMenuItem(v)} />
                </div>
              </Drawer>
              <AppBar position="fixed" className={classes.appBar}>
                <Toolbar>
                  <Typography variant="title" color="inherit" className={classes.flex}>
                    <a className="logo" href="/">Gotown</a>
                  </Typography>
                  <IconButton onClick={this.handleDrawer(true)}>
                    <MenuIcon />
                  </IconButton>
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
