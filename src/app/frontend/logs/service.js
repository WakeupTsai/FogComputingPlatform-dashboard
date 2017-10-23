<<<<<<< HEAD
// Copyright 2017 The Kubernetes Dashboard Authors.
=======
// Copyright 2017 The Kubernetes Authors.
>>>>>>> upstream/master
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Service class for logs management.
 * @final
 */
export class LogsService {
  /**
   * @ngInject
   */
  constructor() {
    /** @private {boolean} */
    this.inverted_ = true;

    /** @private {boolean} */
    this.compact_ = false;

    /** @private {boolean} */
    this.showTimestamp_ = false;
<<<<<<< HEAD
  }

  /**
   * Getter for inverted flag.
   * @return {boolean}
   */
  getInverted() {
    return this.inverted_;
  }

  /**
   * Switches the inverted flag.
   */
=======

    /** @private {boolean} */
    this.previous_ = false;

    /** @private {boolean} */
    this.following_ = false;
  }

  setFollowing() {
    this.following_ = !this.following_;
  }

  /**
   * @return {boolean}
   * @export
   */
  getFollowing() {
    return this.following_;
  }

>>>>>>> upstream/master
  setInverted() {
    this.inverted_ = !this.inverted_;
  }

  /**
<<<<<<< HEAD
   * Switches the compact flag.
   */
=======
   * @return {boolean}
   * @export
   */
  getInverted() {
    return this.inverted_;
  }

>>>>>>> upstream/master
  setCompact() {
    this.compact_ = !this.compact_;
  }

  /**
<<<<<<< HEAD
   * Getter for compact flag.
   * @return {boolean}
=======
   * @return {boolean}
   * @export
>>>>>>> upstream/master
   */
  getCompact() {
    return this.compact_;
  }

<<<<<<< HEAD
  /**
   * Switches the show timestamp flag
   */
=======
>>>>>>> upstream/master
  setShowTimestamp() {
    this.showTimestamp_ = !this.showTimestamp_;
  }

  /**
<<<<<<< HEAD
   * Getter for the show timestamp flag
   * @returns {boolean}
=======
   * @return {boolean}
   * @export
>>>>>>> upstream/master
   */
  getShowTimestamp() {
    return this.showTimestamp_;
  }
<<<<<<< HEAD
=======

  setPrevious() {
    this.previous_ = !this.previous_;
  }

  /**
   * @return {boolean}
   * @export
   */
  getPrevious() {
    return this.previous_;
  }
>>>>>>> upstream/master
}
