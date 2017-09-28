// phstsai

/**
 * Number of labels, that are always visible.
 * @type {number}
 */
const alwaysVisibleLabelsNumber = 5;

/**
 * Regular expression for URL validation created by @dperini.
 * https://gist.github.com/dperini/729294
 *
 * @type {RegExp}
 */
const urlRegexp = new RegExp(
    '^(?:(?:https?|ftp)://)(?:\\S+(?::\\S*)?@)?(?:(?!(?:10|127)(?:\\.\\d{1,3}){3})(?!(?:169\\.254|192\\.168)' +
        '(?:\\.\\d{1,3}){2})(?!172\\.(?:1[6-9]|2\\d|3[0-1])(?:\\.\\d{1,3}){2})(?:[1-9]\\d?|1\\d\\d|2[01]\\d|' +
        '22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[1-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?' +
        ':(?:[a-z\\u00a1-\\uffff0-9]-*)*[a-z\\u00a1-\\uffff0-9]+)(?:\\.(?:[a-z\\u00a1-\\uffff0-9]-*)*[a-z\\u' +
        '00a1-\\uffff0-9]+)*(?:\\.(?:[a-z\\u00a1-\\uffff]{2,}))\\.?)(?::\\d{2,5})?(?:[/?#]\\S*)?$',
    'i');

/**
 * @final
 */
export default class SensorLabelsController {
  /**
   * Constructs labels controller.
   * @ngInject
   */
  constructor() {
    /** @export {!Object<string, string>} Initialized from the scope. */
    this.labels;

    /** @private {boolean} */
    this.isShowingAllLabels_ = false;
  }

  /**
   * Returns true if element at index position should be visible.
   *
   * @param {number} index
   * @return {boolean}
   * @export
   */
  isVisible(index) {
    return this.isShowingAllLabels_ || index < alwaysVisibleLabelsNumber;
  }

  /**
   * Returns true if more labels than alwaysVisibleLabelsNumber is available.
   *
   * @return {boolean}
   * @export
   */
  isMoreAvailable() {
    return Object.keys(this.labels).length > alwaysVisibleLabelsNumber;
  }

  /**
   * Returns true if all labels are showed.
   *
   * @return {boolean}
   * @export
   */
  isShowingAll() {
    return this.isShowingAllLabels_;
  }

  /**
   * Switches labels view between two states, which are showing only alwaysVisibleLabelsNumber of
   * labels and showing all labels.
   *
   * @export
   */
  switchLabelsView() {
    this.isShowingAllLabels_ = !this.isShowingAllLabels_;
  }

  /**
   * Returns true if given value is URL.
   *
   * @return {boolean}
   * @export
   */
  isHref(value) {
    return urlRegexp.test(value.trim());
  }
}

/**
 * Labels component definition.
 *
 * @type {!angular.Component}
 */
/** phstsai **/
export const sensorLabelComponent= {
  bindings: {
    'sensorlabels': '=',
  },
  controller: SensorLabelsController,
  controllerAs: 'sensorLabelsCtrl',
  templateUrl: 'common/components/sensorlabel/sensorlabels.html',

};
