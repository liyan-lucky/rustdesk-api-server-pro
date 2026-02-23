"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = loader;
Object.defineProperty(exports, "defaultMinimizerOptions", {
  enumerable: true,
  get: function () {
    return _utils.defaultMinimizerOptions;
  }
});
var _plugins = require("./plugins");
var _utils = require("./utils");
var _options = _interopRequireDefault(require("./options.json"));
function _interopRequireDefault(e) { return e && e.__esModule ? e : { default: e }; }
async function loader(content) {
  const rawOptions = this.getOptions(_options.default);
  const options = (0, _utils.normalizeOptions)(rawOptions, this);
  if (options.preprocessor) {
    // eslint-disable-next-line no-param-reassign
    content = await options.preprocessor(content, this);
  }
  const plugins = [];
  const errors = [];
  const imports = [];
  const replacements = [];
  let isSupportAbsoluteURL = false;

  // TODO enable by default in the next major release
  if (this._compilation && this._compilation.options && this._compilation.options.experiments && this._compilation.options.experiments.buildHttp) {
    isSupportAbsoluteURL = true;
  }
  if (options.sources) {
    plugins.push((0, _plugins.sourcesPlugin)({
      isSupportAbsoluteURL,
      isSupportDataURL: options.esModule,
      sources: options.sources,
      resourcePath: this.resourcePath,
      context: this.context,
      imports,
      errors,
      replacements
    }));
  }
  if (options.minimize) {
    plugins.push((0, _plugins.minimizerPlugin)({
      minimize: options.minimize,
      errors
    }));
  }
  let {
    html
  } = await (0, _utils.pluginRunner)(plugins).process(content);
  for (const error of errors) {
    this.emitError(error instanceof Error ? error : new Error(error));
  }
  const isTemplateLiteralSupported = (0, _utils.supportTemplateLiteral)(this);
  html = (isTemplateLiteralSupported ? (0, _utils.convertToTemplateLiteral)(html) : JSON.stringify(html)

  // Invalid in JavaScript but valid HTML
  ).replace(/[\u2028\u2029]/g, str => str === "\u2029" ? "\\u2029" : "\\u2028");
  if (options.postprocessor) {
    // eslint-disable-next-line no-param-reassign
    html = await options.postprocessor(html, this);
  }
  const importCode = (0, _utils.getImportCode)(html, imports, options);
  const moduleCode = (0, _utils.getModuleCode)(html, replacements, this, {
    esModule: options.esModule,
    isTemplateLiteralSupported
  });
  const exportCode = (0, _utils.getExportCode)(html, options);
  return `${importCode}${moduleCode}${exportCode}`;
}