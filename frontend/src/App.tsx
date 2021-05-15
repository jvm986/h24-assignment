import Container from "@material-ui/core/Container";
import Box from "@material-ui/core/Box";
import {
  createStyles,
  makeStyles,
  Typography,
  Theme,
  Grid,
  LinearProgress,
} from "@material-ui/core";

import { getUrlInfo, validURL } from "./services/services";
import { UrlEntry } from "./components/UrlEntry";
import { useState } from "react";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      padding: theme.spacing(3, 2),
      display: "flex",
      justifyContent: "center",
    },
  })
);

type urlInfo = {
  url: string;
  title: string;
  htmlVersion: string;
  headings: heading[];
  links: link[];
  login: boolean;
};

type heading = {
  importance: number;
  count: number;
};

type link = {
  link: string;
  internal: boolean;
  accessible: boolean;
};

const App: React.FC = () => {
  const classes = useStyles();
  const [url, setUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [urlInfo, setUrlInfo] = useState<urlInfo>();
  const [validUrl, setValidUrl] = useState(false);
  const [linksChecked, setLinksChecked] = useState(false);
  const [linkProgress, setLinkProgress] = useState(0);
  const [error, setError] = useState("");

  function checkLinks(links: Array<link>): Promise<any> {
    var checked = 0;
    return Promise.all(
      links.map(async (link: link, index) => {
        try {
          const response = await getUrlInfo(link.link);
          link.accessible = response.status == 200;
          checked++;
          setLinkProgress((checked / links.length) * 100);
        } catch (err) {
          checked++;
          setLinkProgress((checked / links.length) * 100);
          return link;
        }
        return link;
      })
    );
  }

  async function handleSubmit() {
    setUrlInfo(undefined);
    setLinkProgress(0);
    setError("");
    setLinksChecked(false);
    setLoading(true);
    try {
      const response = await getUrlInfo(url);
      setUrl(response.data.url);
      setUrlInfo(response.data);
      setLoading(false);
      if (response.data.links) {
        await checkLinks(response.data.links);
      }
    } catch (err) {
      setError(err.toString());
    }
    setLinksChecked(true);
    setLinkProgress(100);
  }

  const handleChange = (url: string) => {
    setLinkProgress(0);
    setError("");
    setLinksChecked(false);
    setUrlInfo(undefined);
    setLoading(false);
    setUrl(url);
    setValidUrl(validURL(url));
  };

  return (
    <Grid container className={classes.root} alignItems="center">
      <Grid item xs={12} md={6} lg={4}>
        <Box>
          <UrlEntry
            url={url}
            handleChange={handleChange}
            handleSubmit={handleSubmit}
            validUrl={validUrl}
            loading={loading}
          />
        </Box>
        {urlInfo ? (
          <Box>
            <Box mt={4} mb={2}>
              <LinearProgress variant="determinate" value={linkProgress} />
            </Box>
            <Typography variant="h5">
              {urlInfo.title.replace("&amp;", "&")}
            </Typography>
            <Typography>HTML Version: {urlInfo.htmlVersion}</Typography>
            <Typography>
              {urlInfo.login
                ? "Has a login form"
                : "Does not have a login form"}
            </Typography>
            <Typography>
              {urlInfo.links
                ? urlInfo.links.filter((link) => !link.internal).length
                : 0}{" "}
              External Links
              {linksChecked
                ? urlInfo.links
                  ? " : " +
                    urlInfo.links.filter(
                      (link) => link.accessible && !link.internal
                    ).length +
                    " accessible"
                  : ""
                : ": checking accessibility..."}
            </Typography>
            <Typography>
              {urlInfo.links
                ? urlInfo.links.filter((link) => link.internal).length
                : 0}{" "}
              Internal Links
              {linksChecked
                ? urlInfo.links
                  ? " : " +
                    urlInfo.links.filter(
                      (link) => link.accessible && link.internal
                    ).length +
                    " accessible"
                  : ""
                : " : checking accessibility..."}
            </Typography>
            <Typography variant="h1">
              {urlInfo.headings[0].count} H1 Tags
            </Typography>
            <Typography variant="h2">
              {urlInfo.headings[1].count} H2 Tags
            </Typography>
            <Typography variant="h3">
              {urlInfo.headings[2].count} H3 Tags
            </Typography>
            <Typography variant="h4">
              {urlInfo.headings[3].count} H4 Tags
            </Typography>
            <Typography variant="h5">
              {urlInfo.headings[4].count} H5 Tags
            </Typography>
            <Typography variant="h6">
              {urlInfo.headings[5].count} H6 Tags
            </Typography>
          </Box>
        ) : error ? (
          <Box mt={5}>
            <Typography>{error}</Typography>
          </Box>
        ) : (
          ""
        )}
      </Grid>
    </Grid>
  );
};

export default App;
